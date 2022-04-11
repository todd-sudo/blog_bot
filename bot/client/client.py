import aiohttp

from config.config import get_domain
from logger.logger import logger


domain = get_domain()


class Client:
    headers: dict
    cookies: dict
    url: str

    def __init__(
            self, headers: dict = None,
            cookies: dict = None,
            url: str = None
    ):
        self.headers = headers
        self.cookies = cookies
        self.url = domain + url

    async def get(self):
        async with aiohttp.ClientSession() as session:
            response = await session.get(
                url=self.url, headers=self.headers
            )
            return response

    async def post(self, data: dict):
        async with aiohttp.ClientSession() as session:
            response = await session.post(
                    url=self.url, json=data, headers=self.headers
            )
            return response

    async def delete(self):
        async with aiohttp.ClientSession() as session:
            response = await session.delete(url=self.url, headers=self.headers)
            return response
