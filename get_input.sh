#!/bin/bash
echo $date_input
echo $team_input

bq query --nouse_legacy_sql --parameter="mys:string:${date_input}" --parameter="myt:string:${team_input}" 'CREATE OR REPLACE TABLE `data-ingest-416321.ncaab_mens.input` 
AS SELECT
    game_date,
    team_name,
    team_winner,
    total_rebounds,
    field_goals_attempted,
    free_throws_attempted,
    steals
FROM `data-ingest-416321.ncaab_mens.subset2024`
WHERE game_date = CAST(@mys AS DATE) AND team_name = @myt
'
