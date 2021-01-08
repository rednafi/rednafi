<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>
    
```python
​
import json
from dataclasses import asdict, dataclass


@dataclass
class Stack:
    languages   : tuple[str, ...] = ("Python", "Bash")
    databases   : tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : tuple[str, ...] = ("Docker", "Celery", "RQ")
    ongoing     : tuple[str, ...] = ("Django", "DRF", "JavaScript")

    def serialize(self):
        return json.dumps(asdict(self), indent=4)


stack = Stack()
print(stack.serialize())
​
```
</h3>
