from aiogram import types
from aiogram.dispatcher import FSMContext

from client.client import Client
from config.loader import dp
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


@dp.message_handler(commands=["all_categories"])
async def get_all_categories(message: types.Message):
    headers = {
        "user_tg_id": str(message.from_user.id)
    }
    client = Client(url="/api/categories/", headers=headers)
    response = await client.get()
    data = await response.json()
    print(data)


