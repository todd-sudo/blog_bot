import logging

from aiogram import executor

from config.loader import bot, dp
import handlers

# async def scheduler():
#     aioschedule.every(4).hours.do(check_birthday)
#     while True:
#         await aioschedule.run_pending()
#         await asyncio.sleep(1)


async def on_startup(_):
    # asyncio.create_task(scheduler())
    await bot.delete_webhook()


async def on_shutdown(dp):
    logging.warning("Shutting down..")
    await bot.delete_webhook()
    await dp.storage.close()
    await dp.storage.wait_closed()
    logging.warning("Bot down")


if __name__ == '__main__':
    executor.start_polling(
        dp,
        on_startup=on_startup,
        on_shutdown=on_shutdown,
        skip_updates=True
    )
