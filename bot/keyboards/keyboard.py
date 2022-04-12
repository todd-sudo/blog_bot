from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton
from aiogram.utils.callback_data import CallbackData


def profile_keyboards() -> InlineKeyboardMarkup:
    keyboard = InlineKeyboardMarkup()
    category = InlineKeyboardButton(
        "Категории",
        callback_data='categories'
    )
    posts = InlineKeyboardButton(
        "Записи",
        callback_data="posts"
    )
    keyboard.add(category)
    keyboard.add(posts)
    return keyboard


def create_posts_keyboards():
    post_keyboard = InlineKeyboardMarkup()
    add_post = InlineKeyboardButton(
        "Создать запись",
        callback_data="add_post"
    )
    post_keyboard.add(add_post)
    return post_keyboard


def create_category_keyboards():
    post_keyboard = InlineKeyboardMarkup()
    add_post = InlineKeyboardButton(
        "Создать категорию",
        callback_data="add_category"
    )
    post_keyboard.add(add_post)
    return post_keyboard


def get_callback_data_category():
    cb = CallbackData("c", "c_id")
    return cb


def get_callback_data_post():
    cb = CallbackData("post", "post_id")
    return cb

