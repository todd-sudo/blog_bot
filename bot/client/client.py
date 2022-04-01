import aiohttp

from config.config import get_domain
from logger.logger import logger


domain = get_domain()


class Client:
    headers: dict
    cookies: dict

    def __init__(self, headers: dict = None, cookies: dict = None):
        self.headers = headers
        self.cookies = cookies

    async def get(self):
        async with aiohttp.ClientSession() as session:
            pass

    async def post(self, data: dict):
        url = domain + "/api/user/create-user"
        async with aiohttp.ClientSession() as session:
            response = await session.post(
                    url=url, json=data, headers=self.headers
            )
            return response

    async def delete(self):
        async with aiohttp.ClientSession() as session:
            pass
