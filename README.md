<!-- Zero width character is used to put extra blank lines before and after code -->

<h3>
    
```python
​
from __future__ import annotations

import json
from dataclasses import asdict, dataclass


@dataclass
class Arsenal:
    languages: tuple[str, ...] = ("Python", "JS", "Go")
    databases: tuple[str, ...] = ("SQLite", "PostgreSQL", "DynamoDB", "Redis")
    misc     : tuple[str, ...] = ("Docker", "Celery", "RabbitMQ", "Arq", "SQS")
    ongoing  : tuple[str, ...] = ("Django", "DRF", "Asyncio")

    def jsonify(self) -> None:
        return json.dumps(asdict(self), indent=4)


arsenal = Arsenal()
print(arsenal.jsonify())
​
```
</h3>
