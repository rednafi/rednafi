### Hi there 👋

<!--
**rednafi/rednafi** is a ✨ _special_ ✨ repository because its `README.md` (this file) appears on your GitHub profile.

Here are some ideas to get you started:

- 🔭 I’m currently working on ...
- 🌱 I’m currently learning ...
- 👯 I’m looking to collaborate on ...
- 🤔 I’m looking for help with ...
- 💬 Ask me about ...
- 📫 How to reach me: ...
- 😄 Pronouns: ...
- ⚡ Fun fact: ...
-->

```python
from dataclasses import dataclass
from typing import Tuple

@dataclass
class Bio:
    name        : str = "Redowan Delowar"
    designation : str = "Software Engineer"
    workplace   : str = "DendiSoftware, Raleigh, North Carolina"
    base        : str = "Dhaka, Bangladesh"
    blog        : str = "rednafi.github.io/digressions"

@dataclass
class Stack:
    languages   : Tuple[str, ...] = ("Python", "Go", "Shell")
    databases   : Tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : Tuple[str, ...] = ("Docker", "Celery")
    ongoing     : Tuple[str, ...] = ("Django", "GraphQL")

@dataclass
class Social:
    twitter     : str = "rednafi"
    linkedin    : str = "redowan"
```
