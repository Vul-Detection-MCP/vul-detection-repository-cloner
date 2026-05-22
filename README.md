# repository-cloner

## Description

This tool provides a way to clone list of Git repositories of given commits using a CSV file.

Use these volumes:

- `/app/list.csv` - source.
- `/app/git` - cloned repositories.

Unique ID is needed to distinguish multiple clones of the same project. 
You may choose any convenient strategy for value generation.
