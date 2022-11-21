#!/bin/bash

# Legacy & Transaction Script & Active Record

curl http://localhost:3000/pets
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "big", "age": 1}'
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "alpha", "age": 2}'
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "shiro", "age": 3}'
curl http://localhost:3000/pets/3eeac5e1-aacf-445f-8f6e-805385d0d5e7
curl -X POST http://localhost:3000/pets/3eeac5e1-aacf-445f-8f6e-805385d0d5e7/change-name -H 'Content-Type: application/json' -d '{"name": "big2"}'
curl -X POST http://localhost:3000/pets/3eeac5e1-aacf-445f-8f6e-805385d0d5e7/sell
curl -X POST http://localhost:3000/pets/3eeac5e1-aacf-445f-8f6e-805385d0d5e7/return
curl http://localhost:3000/pets/98e03cd9-16f9-4dac-932c-f6d52dda7126
curl -X POST http://localhost:3000/pets/98e03cd9-16f9-4dac-932c-f6d52dda7126/change-name -H 'Content-Type: application/json' -d '{"name": "shiroro"}'
curl -X POST http://localhost:3000/pets/98e03cd9-16f9-4dac-932c-f6d52dda7126/sell
curl -X POST http://localhost:3000/pets/98e03cd9-16f9-4dac-932c-f6d52dda7126/return

# Domain Model

curl http://localhost:3000/pets
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "bang", "age": 1}'
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "bong", "age": 2}'
curl -X POST http://localhost:3000/pets/56ea8037-071f-4e3b-9898-de188c91e6c3/change-name -H 'Content-Type: application/json' -d '{"name": "bang2"}'
curl -X POST http://localhost:3000/pets/f985f6f4-1eab-4a81-8294-8761899fbf3f/sell
curl -X POST http://localhost:3000/pets/f985f6f4-1eab-4a81-8294-8761899fbf3f/return


# Event Sourced Domain Model

curl http://localhost:3000/pets
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "elon", "age": 20}'
curl -X POST http://localhost:3000/pets -H 'Content-Type: application/json' -d '{"name": "elon2", "age": 22}'
curl -X POST http://localhost:3000/pets/1ac2b96a-6199-4c3b-86d5-d8bf8068569d/change-name -H 'Content-Type: application/json' -d '{"name": "elonmusk"}'
curl -X POST http://localhost:3000/pets/1ac2b96a-6199-4c3b-86d5-d8bf8068569d/sell
curl -X POST http://localhost:3000/pets/1ac2b96a-6199-4c3b-86d5-d8bf8068569d/return
curl -X POST http://localhost:3000/pets/8c7d5676-a411-4c69-ac39-eff508ea834b/sell
curl -X POST http://localhost:3000/pets/8c7d5676-a411-4c69-ac39-eff508ea834b/return
curl -X POST http://localhost:3000/pets/8c7d5676-a411-4c69-ac39-eff508ea834b/change-name -H 'Content-Type: application/json' -d '{"name": "geohit"}'






