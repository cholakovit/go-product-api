updated_at doesn't change when update product

messages doesn't change when the error is different. Example:
title product is missing - got an error msg, and when title is not missing, but it is less then 2 char, i got the same error msg
for missing title. I have to refresh the server to get the new error

category, desc less then 20 char doesn't show error PATCH AND POST
category, empty desc doesn't show error             PATCH AND POST

on users.model.go, if i remove * from *string required, i have an error

i register token, and on the next day i still have it, for users