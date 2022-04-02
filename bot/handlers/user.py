import json

from aiogram import types

from client.client import Client
from config.loader import dp
from keyboards.category import get_all_categories_keyboards


@dp.message_handler(commands=["start"])
async def create_user(message: types.Message):
    client = Client(url="/api/user/create-user")
    response = await client.post(data={
        "first_name": message.from_user.first_name,
        "last_name": message.from_user.last_name,
        "username": message.from_user.username,
        "user_tg_id": message.from_user.id
    })
    if response.status in [200, 201]:
        await message.answer("Вы успешно зарегистрированы!")
    elif response.status == 409:
        await message.answer("Вы уже зарегистрированы!")
    else:
        await message.answer("Ошибка! Обратитесь к администратору! /admin")


@dp.message_handler(commands=["profile"])
async def profile_user(message: types.Message):
    headers = {
        "user_tg_id": str(message.from_user.id)
    }
    client = Client(headers=headers, url="/api/user/profile")
    response = await client.get()
    if response.status == 404:
        await message.answer("Такого пользователя нет!")
    elif response.status == 409:
        await message.answer("Неизвестная ошибка! Обратитесь к администратору!")
    else:
        data = await response.json()
        data = data.get("data")
        username = data.get("username")
        first_name = "Нет" if not data.get('first_name') else data.get('first_name')
        last_name = "Нет" if not data.get('last_name') else data.get('last_name')
        user_tg_id = data.get("user_tg_id")
        categories = data.get("categories")
        print(categories)
        await message.answer(
            f"Ваш профиль:\n\nUsername - {username}\n"
            f"Имя - {first_name}\nФамилия - {last_name}\n"
            f"Телеграм ID - {user_tg_id}",
            reply_markup=get_all_categories_keyboards()
        )
