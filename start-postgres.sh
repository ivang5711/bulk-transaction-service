#!/bin/bash

sudo docker compose exec -it db psql -U bts -d btsdb
