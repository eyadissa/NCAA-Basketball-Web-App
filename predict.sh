#!/bin/bash
bq query --apilog=stdout --nouse_legacy_sql \ '
SELECT predicted_team_winner
FROM ML.PREDICT(MODEL `ncaab_mens.wins_logistic4`,
   (
     SELECT
      total_rebounds,
      field_goals_attempted,
      free_throws_attempted,
      steals
FROM
 `data-ingest-416321.ncaab_mens.input`
   )
)'

bq query --apilog=stdout --nouse_legacy_sql \ '
SELECT
    team_winner
FROM `data-ingest-416321.ncaab_mens.input`
'
