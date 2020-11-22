# Example

Suppose that you want to create a Secret Santa for some members of the
fellowship of the Ring as well as some significant elves. Of course, some
members of the fellowship might not want to be each other's Secret Santa
because they might be buying gifts for each other outside of the Secret
Santa because they are so close.

Here's an example input file (in YAML format).

```yaml
- name: Gandalf
  invalid:
    - Elrond
    - Galadriel
- name: Elrond
  invalid:
    - Gandalf
    - Galadriel
- name: Galadriel
  invalid:
    - Gandalf
    - Elrond
- name: Aragorn
  invalid:
    - Arwen
- name: Arwen
  invalid:
    - Aragorn
- name: Legolas
  invalid:
    - Gimli
- name: Gimli
  invalid:
    - Legolas
- name: Frodo
  invalid:
    - Samwise
- name: Samwise
  invalid:
    - Frodo
- name: Pippin
- name: Merry
```

To determine who is bying for how, simply execute the `simple-santa` program.
Here is what you might get:

```sh
$ simple-santa fellowship.yaml
Gimli     ==> Elrond
Merry     ==> Galadriel
Gandalf   ==> Aragorn
Galadriel ==> Legolas
Aragorn   ==> Gandalf
Frodo     ==> Arwen
Samwise   ==> Gimli
Pippin    ==> Samwise
Elrond    ==> Pippin
Arwen     ==> Merry
Legolas   ==> Frodo
```

Each time you run the results will be different, so don't expect to get the
above results.
