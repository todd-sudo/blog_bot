from aiogram import types
from aiogram.dispatcher import filters

from client.client import Client
from config.loader import dp


@dp.message_handler(filters.Text(contains="оллар", ignore_case=True))
async def test(message: types.Message):
    await message.answer("egijegj")


@dp.message_handler(commands=["start"])
async def create_user(message: types.Message):
    client = Client()
    response = await client.post(data={
        "first_name": message.from_user.first_name,
        "last_name": message.from_user.last_name,
        "username": message.from_user.username,
        "user_tg_id": message.from_user.id
    })
    if response.status == 200:
        await message.answer("Вы успешно зарегистрированы!")
    elif response.status == 409:
        await message.answer("Вы уже зарегистрированы!")
    else:
        await message.answer("Ошибка! Обратитесь к администратору! /admin")

