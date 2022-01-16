#!/bin/bash

awslocal dynamodb create-table --cli-input-json file:///init/data/dynamodb/todo.json

awslocal dynamodb list-tables
