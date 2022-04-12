from aiogram import types
from aiogram.dispatcher import FSMContext

from client.client import Client
from config.loader import dp, bot
from keyboards.keyboard import (
    create_posts_keyboards,
    get_callback_data_category,
    get_callback_data_post
)
from states.create_post import PostState


cb_post = get_callback_data_post()
cb_c = get_callback_data_category()


@dp.callback_query_handler(text="posts")
async def get_all_post(call: types.CallbackQuery):

    headers = {
        "user_tg_id": str(call.from_user.id)
    }
    client = Client(url="/api/posts/", headers=headers)
    response = await client.get()
    res_json = await response.json()
    if response.status == 200:
        posts = res_json.get("data")
        if not posts:
            await call.message.answer(
                "У вас нет записей ✖", reply_markup=create_posts_keyboards()
            )
        else:
            for post in posts:
                button = types.InlineKeyboardButton(
                    text="Удалить ❌",
                    callback_data=cb_post.new(post_id=post.get("id"))
                )

                reply_markup = types.InlineKeyboardMarkup()
                reply_markup.add(button)

                title_post = post.get("title")
                content_post = post.get("content")
                category_post = post.get("category").get("name")

                await call.message.answer(
                    f"➕ {title_post}\n➕ {category_post}\n➕ {content_post}",
                    reply_markup=reply_markup
                )
    else:
        await call.message.answer("Ошибка! Обратитесь к администратору! /admin")


@dp.callback_query_handler(text="add_post")
async def add_post(call: types.CallbackQuery):
    await call.message.answer("Введите название записи...")
    await PostState.name.set()


@dp.message_handler(state=PostState.name)
async def set_name_post(message: types.Message, state: FSMContext):
    await state.update_data(c_name=message.text)
    await message.answer("Введите текст записи...")
    await PostState.content.set()


@dp.message_handler(state=PostState.content)
async def set_content_post(message: types.Message, state: FSMContext):
    await state.update_data(c_content=message.text)
    headers = {
        "user_tg_id": str(message.from_user.id)
    }
    client = Client(url="/api/categories/", headers=headers)
    response = await client.get()
    data = await response.json()
    categories = data.get("data")
    if categories:
        for c in categories:
            button = types.InlineKeyboardButton(
                text=c.get("name"),
                callback_data=cb_c.new(c_id=c.get("id"))
            )

            reply_markup = types.InlineKeyboardMarkup()
            reply_markup.add(button)
        await message.answer("Выберите категорию", reply_markup=reply_markup)
        await PostState.category.set()

    else:
        await message.answer(
            "У вас нет категорий. Вы не можете создать запись"
        )
        await state.reset_state()


@dp.callback_query_handler(cb_c.filter(), state=PostState.category)
async def set_category_post(
        call: types.CallbackQuery, callback_data: dict, state: FSMContext
):
    category_id = callback_data["c_id"]
    category_data = await state.get_data()

    headers = {
        "user_tg_id": str(call.message.from_user.id)
    }
    request_data = {
        "title": category_data.get("c_name"),
        "content": category_data.get("c_content"),
        "user_tg_id": call.from_user.id,
        "category_id": int(category_id)
    }
    client = Client(url="/api/posts/", headers=headers)
    response = await client.post(data=request_data)

    if response.status == 201:
        await call.message.answer("Запись успешно добавлена 👌")
    else:
        await call.message.answer(
            "Произошла ошибка. Обратитесь к администратору /admin"
        )
    await state.finish()


@dp.callback_query_handler(cb_post.filter())
async def delete_post(call: types.CallbackQuery, callback_data: dict):
    post_id = callback_data.get("post_id")
    headers = {
        "user_tg_id": str(call.from_user.id)
    }
    client = Client(url=f"/api/posts/{post_id}", headers=headers)
    response = await client.delete()
    print(response.status)
    if response.status == 204:
        msg = call.message.message_id
        await bot.delete_message(call.message.chat.id, msg)
        await call.message.answer("Запись успешно удалена 🆗")
    else:
        await call.message.answer(
            "Ошибка! Обратитесь к администратору 🙊 /admin"
        )

