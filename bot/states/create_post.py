from aiogram.dispatcher.filters.state import StatesGroup, State


class PostState(StatesGroup):
    name = State()
    content = State()
    category = State()
