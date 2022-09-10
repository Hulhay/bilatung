import os
import faker
import dotenv
import requests as r
from assertpy import assert_that

dotenv.load_dotenv()
base_url = os.environ['BASE_URL_TESTING']
fake = faker.Faker()

def test_register_empty_body():
    resp = r.post(f'{base_url}/register')
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('body in body is required')

def test_register_empty_email():
    body = {
        "username": "thisIsUsername",
        "password": "thisIsPassword"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()
    print(data)

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('email in body is required')

def test_register_empty_password():
    body = {
        "email": "this.is.email@gmail.com"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('password in body is required')

def test_register_empty_username():
    body = {
        "email": "this.is.email@gmail.com",
        "password": "thisIsPassword"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('username in body is required')

def test_register_invalid_email_format():
    body = {
        "email": "thisIsInvalidEmailFormat",
        "password": "thisIsPassword",
        "username": "thisIsUname"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).contains('email in body must be of type email')

def test_register_invalid_username():
    body = {
        "email": "this.is.email@gmail.com",
        "password": "thisIsPassword",
        "username": "u"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('username in body should be at least 5 chars long')

def test_register_invalid_password():
    body = {
        "email": "this.is.email@gmail.com",
        "password": "p",
        "username": "thisIsUname"
    }

    resp = r.post(f'{base_url}/register', json=body)
    data = resp.json()

    assert_that(resp.status_code).is_not_equal_to(200)
    assert_that(data).contains('message')
    assert_that(data['message']).is_equal_to('password in body should be at least 6 chars long')

def test_register_success():

    username = f'fake_{fake.name().replace(" ", "")}'
    email = f'{username}@gmail.com'

    body = {
        "email": username,
        "password": "thisIsFakePassword",
        "username": email,
    }

    resp = r.post(f'{base_url}/register', json=body)

    assert_that(resp.status_code).is_equal_to(201)