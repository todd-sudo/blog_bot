import os
from dotenv import load_dotenv


dotenv_path = os.path.join(os.path.dirname(".."), '.env')
if os.path.exists(dotenv_path):
    load_dotenv(dotenv_path)


def get_domain():
    return "http://127.0.0.1:8000"


def get_token_bot() -> str:
    """ Возвращает токен Телеграм бота
    """
    return os.getenv("TOKEN")
