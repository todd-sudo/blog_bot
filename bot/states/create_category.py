from aiogram.dispatcher.filters.state import StatesGroup, State


class CategoryState(StatesGroup):
    name = State()
