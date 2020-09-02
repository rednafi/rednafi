
<!-- service: https://carbon.now.sh/?bg=rgba(171%2C%20184%2C%20195%2C%201)&t=a11y-dark&wt=none&l=python&ds=true&dsyoff=20px&dsblur=68px&wc=true&wa=false&pv=56px&ph=56px&ln=false&fl=1&fm=JetBrains%20Mono&fs=16px&lh=133%25&si=false&es=2x&wm=false&code=from%2520dataclasses%2520import%2520dataclass%250Afrom%2520typing%2520import%2520Tuple%250A%250A%2540dataclass%250Aclass%2520Stack%253A%250A%2520%2520%2520%2520languages%2520%2520%2520%253A%2520Tuple%255Bstr%252C%2520...%255D%2520%253D%2520(%2522Python%2522%252C%2520%2522Go%2522%252C%2520%2522Bash%2522)%250A%2520%2520%2520%2520databases%2520%2520%2520%253A%2520Tuple%255Bstr%252C%2520...%255D%2520%253D%2520(%2522PostgreSQL%2522%252C%2520%2522Mongo%2522%252C%2520%2522Redis%2522)%250A%2520%2520%2520%2520misc%2520%2520%2520%2520%2520%2520%2520%2520%253A%2520Tuple%255Bstr%252C%2520...%255D%2520%253D%2520(%2522Docker%2522%252C%2520%2522Celery%2522)%250A%2520%2520%2520%2520ongoing%2520%2520%2520%2520%2520%253A%2520Tuple%255Bstr%252C%2520...%255D%2520%253D%2520(%2522Django%2522%252C%2520%2522GraphQL%2522)
width: Fixed, 950
theme: A11yDark

from dataclasses import dataclass
from typing import Tuple

@dataclass
class Stack:
    languages   : Tuple[str, ...] = ("Python", "Go", "Bash")
    databases   : Tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : Tuple[str, ...] = ("Docker", "Celery")
    ongoing     : Tuple[str, ...] = ("Django", "GraphQL")
 -->

![carbon](https://user-images.githubusercontent.com/30027932/92017490-ae553e00-ed75-11ea-9419-bd13d04efe4b.png)
