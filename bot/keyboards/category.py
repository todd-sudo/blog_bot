from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton
from aiogram.utils.callback_data import CallbackData


def get_all_categories_keyboards() -> InlineKeyboardMarkup:

    delete_message_keyboard = InlineKeyboardMarkup()
    delete = InlineKeyboardButton(
        "Категории",
        callback_data='categories'
    )
    delete_message_keyboard.add(delete)
    return delete_message_keyboard


def get_callback_data():
    cb = CallbackData("id", "c_id")
    return cb

