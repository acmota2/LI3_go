# LI3_go

This repository was used to learn the programming language Go. As a template I used the project of the _Laboratórios de Informática 3_ subject from the _Licenciatura em Engenharia Informática_ of Universidade do Minho from the year 2021/22.

This is the [statement](./misc/guiao2.pdf), which unfortunately I can only provide in Portuguese.

This was originally a project with the objective to train algorithms and data structures, recurring to the C programming language and the **glib** external library.

## Preamble

A repository consisting of 3 .csv files was provided with the project's statement, but given the file limit size of GitHub, unfortunately I could not upload it, having just the short tests inside the [tests](./tests) folder provided here.

As it was with the original project, as I used the exact same files, error correction or recuperation was not expected, nor evaluated, maybe with the exception of file opening errors and such.

There are 3 .csv files, all of them correlated, which's fields and types are as follows:

* `commits.csv`:
    * repo_id - integer
    * author_id - integer
    * committer_id - integer
    * commit_at - date (YYYY-MM-DD hh:mm:ss)
    * message - string

* `repos.csv`:
    * id - integer
    * owner_id - integer
    * full_name - string
    * license - string
    * has_wiki - bool true/false
    * description - string
    * language - string
    * default_branch - string
    * created_at - date (YYYY-MM-DD hh:mm:ss)
    * updated_at - date (YYYY-MM-DD hh:mm:ss)
    * forks_count - integer
    * open_issues - integer
    * stargazers_count - integer
    * size - integer


* `users.csv`:
    * id - integer
    * login - string of characters
    * type - one of the option User/Organization/Bot
    * created_at - date (YYYY-MM-DD hh:mm:ss)
    * followers - integer
    * follower_list - list of integers
    * following - integer
    * following_list - list of integers
    * public_gists - integer
    * public_repos - integer

## Architecture

As a means to an end, it was expected to find a suitable way to store these files within memory, in such a way they could be easily accessible to perform queries over them.

The chosen architecture, very keen on the C architecture I chose when attending this subject, consists of cataloging all 3 files in a slice, which has a view for better access (array of pointers) and a hashtable of references to each element of the catalog, using a suitable id.

A better architecture could have probably be chosen, something more akin to a relational's database internal structure, for example, but as I was using this project as a learning experience, I decided to go for the same kind of idea I went when idealizing it in C.

### Queries

The project initially had 10 queries, of which I chose the first 9 to implement, as the 10th query had already questions regarding how it should be executed at the time of attendance of this subject, and these are as follows:

#### Statistical queries:

* Query 1: present the total of each kind of user
* Query 2: the average number of colaborators in each repository
* Query 3: the number of repositories with bots as contributors
* Query 4: the average number of commits per user

#### Parameterized queries:

All dates from here on are on the format YYYY-MM-DD

* Query 5: given a number **n**, a starting and ending dates, provide the **n** most active users between that date interval
* Query 6: given a number **n** and a programming language's name, provide the top users present in repositories using said programming language
* Query 7: given a date, provide the repositories that were inactive (i.e. didn't have any commit) after said date
* Query 8: given a number **n** and a date, provide the **n** most utilized programming languages from that date onwards
* Query 9: given a number **n**, provide the **n** users which contributed the most to repositories which's owner is their friend, i.e. is present in both their followers and following lists