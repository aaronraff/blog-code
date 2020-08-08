import os


def div(x: int, y: int) -> int:
    return x / y


def get_base_url() -> str:
    return os.getenv('BASE_URL')


def get_user_url(username: str) -> str:
    return f'https://{get_base_url()}/{username}'
