```python
from dataclasses import dataclass
from typing import Tuple

@dataclass
class Stack:
    languages   : Tuple[str, ...] = ("Python", "Go", "Bash")
    databases   : Tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : Tuple[str, ...] = ("Docker", "Celery")
    ongoing     : Tuple[str, ...] = ("Django", "GraphQL")

@dataclass
class Social:
    twitter     : str = "rednafi"
    linkedin    : str = "redowan"
```
