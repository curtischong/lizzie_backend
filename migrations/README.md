How to handle time
- For tables that are frequently used:
  - Store unix time as `unixt`
- For other tables:
  - Store unix time AND local ts as `unixt` and `ts`
  - local ts helps me do analysis
  - unix time helps me convert to whatever I want
  - Why not UTC?
    - Doesn't help with analysis too much
    - Isn't as universal as unix time
-  Store `unixt` up to millisecond precision as a `bigInt` data type

 Table naming convention:
  - Use singular table names. Many reasons on SO
  - Use singular column names.
  - Don't use camelCase use snake_case bc we don't want accidents with double quotes