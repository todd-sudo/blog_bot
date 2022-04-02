from aiogram import types
from aiogram.dispatcher import FSMContext
from aiogram.utils.callback_data import CallbackData

from client.client import Client
from config.loader import dp
from keyboards.category import get_callback_data
from states.create_category import CategoryState


@dp.message_handler(commands=["create_category"])
async def create_category(message: types.Message):
    await message.answer('Введите название категории...')
    await CategoryState.name.set()


@dp.message_handler(state=CategoryState.name)
async def set_name_category(message: types.Message, state: FSMContext):
    category_name = message.text

    client = Client(url="/api/categories/")
    response = await client.post(data={
        "name": category_name,
        "user_tg_id": message.from_user.id
    })
    if response.status == 201:
        await message.answer(f"Категория - {category_name} успешно создана! 😊")
    elif response.status == 400:
        await message.answer("Введите корректные данные!")
    else:
        await message.answer("Ошибка! Обратитесь к администратору! /admin")

    await state.finish()


cb = get_callback_data()


@dp.callback_query_handler(text="categories")
# @dp.message_handler(commands=["all_categories"])
async def get_all_categories(call: types.CallbackQuery):
    headers = {
        "user_tg_id": str(call.from_user.id)
    }
    client = Client(url="/api/categories/", headers=headers)
    response = await client.get()
    data = await response.json()
    categories = data.get("data")
    print(call.data)

    for c in categories:
        button = types.InlineKeyboardButton(
            text=c.get("name"),
            callback_data=cb.new(c_id=c.get("id"))
        )
        reply_markup = types.InlineKeyboardMarkup()
        reply_markup.add(button)
        await call.message.answer(c.get("name"), reply_markup=reply_markup)


@dp.callback_query_handler(cb.filter())
async def callbacks(call: types.CallbackQuery, callback_data: dict):
    _id = callback_data["c_id"]
    print(_id)
