1. For running the project use the command (Git Bash):
 - gin --appPort 3000 --immediate


2. To test the Creating the product use:
{
  "title": "new test product",
  "img": "testImage.webp",
  "desc": "lorem ipsum desc, lorem ipsum desc, lorem ipsum desc",
  "short_desc": "lorem ipsum short desc",
  "price": 10,
  "discount": 2,
  "stock": 100,
  "active": true,
  "manufacturer": "test Manufacturer",
  "thumbs": {
    "thumb1": "test.webp",
    "thumb2": "test.webp",
    "thumb3": "test.webp",
    "thumb4": "test.webp",
    "thumb5": "test.webp"
  },
  "user_id": "1",
  "category_id": "1"
}

3. To test creating user use:
{
  "full_name": "spas cholakov",
  "pass": "1234",
  "email": "aaa@abv.bg",
  "phone": "123123",
  "role": "user",
  "token": "",
  "rtoken": "",
  "created_at": "2023-07-21T12:34:56Z",
  "updated_at": "2023-07-21T12:34:56Z"
}

4. To test creating category use:
  {
    "title": "category 2",
    "desc": "lorem ipsum 2",
    "tags": "lorem, ipsum 2",
    "user_id": "1",
    "category_id": "1"
  }