<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>
    
```python
​
import json
from dataclasses import asdict, dataclass


@dataclass
class Stack:
    languages   : tuple = ("Python", "Bash")
    databases   : tuple = ("PostgreSQL", "Mongo", "Redis")
    misc        : tuple = ("Docker", "Celery")
    ongoing     : tuple = ("Django", "GraphQL", "JavaScript")

    def serialize(self):
        return json.dumps(asdict(self), indent=4)


stack = Stack()
print(stack.serialize())
​
```
</h3>
