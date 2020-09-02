"""
Created using carbon.now.sh
Config:

Theme   : VS Code Material
Width   : Fixed, 850px
Font    : JetBrains Mono
"""

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
    languages   : Tuple[str, ...] = ("Python", "Go", "Bash")
    databases   : Tuple[str, ...] = ("PostgreSQL", "Mongo", "Redis")
    misc        : Tuple[str, ...] = ("Docker", "Celery")
    ongoing     : Tuple[str, ...] = ("Django", "GraphQL")

@dataclass
class Social:
    twitter     : str = "rednafi"
    linkedin    : str = "redowan"
