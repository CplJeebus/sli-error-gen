# SLI Error Gen

Silly script to create error data for sample SLI\SLO graphs. Basic idea is that it will emit as either json or csv output, A set of errors over 180 days. Each error object is a date and number of error mins burned to that date.

## Todo

- Get output working as a csv and json -- `json` DONE
- Create Fast Burn _Bart Simpson_ event(s) -- DONE
- Create Slow Burn events.  -- DONE
- Parameterise all of it. -- DONE
- Add graphing function. -- DONE
- In addition to the running tally of mins burn should include errors burned that day. -- DONE
- Add rampup toggle flag to conf file
- Serious refactoring required 
  - All the variables are named awful
  - Hard to read and clunky 
  - Also some general refactoring, should be able to apply different SLOs to same event. 
