from os import environ

from redis import Redis
from redis.commands.search.field import TextField
from redis.commands.search.indexDefinition import IndexDefinition, IndexType
from redis.commands.json.path import Path
from redis.commands.search.query import Query

from api.models.events import Event

REDIS_HOST = environ.get('REDIS_HOST', '127.0.0.1')
REDIS_PORT = int(environ.get('REDIS_PORT', 6379))


r = Redis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)
    

    
if __name__ == '__main__':

    r = Redis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)
    
    r.zadd('events', )