<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>
    
```python
​
from dataclasses import dataclass, asdict
import json

@dataclass
class Point:
    languages   : tuple = ("Python", "Bash")
    databases   : tuple = ("PostgreSQL", "Mongo", "Redis")
    misc        : tuple = ("Docker", "Celery")
    ongoing     : tuple = ("Django", "GraphQL", "JavaScript")

    def serialize(self):
        stack_dict = asdict(self)
        return json.dumps(stack_dict, indent=4)

point = Point()
print(point.serialize())
​
```
</h3>
