# Secret Santa

"Secret Santa" is a gift-giving game wherin each participant (the giver) is
assigned the name of another participant (the reciever), and the giver
purchases a gift for the reciever. Each person gets a gift! For large families
or work settings, it is a good way to share gifts without making people feel
as though they must spend a lot of time and money purchasing gifts for
everyone.

Traditionally, at some point ahead of the gathering, names are placed into a
hat and each participant draws for whom they will be purchasing. However,
there are some times when getting everyone together is not feasable (for
example, in a global pandemic), so this program serves as a virtual "hat".

To be fair - this program does not generate secret results. At least the
person executing the program will see all the assignments. Maybe a future
enhancement?

# Installation

The simplest method is to go to
https://github.com/SethMMorton/secret-santa/releases/ and download the
binary for your OS. On MacOS/Linux you may have to make the file executable
by executing `chmod +x` on the file.

If you have `go` installed on your computer, you can just use `go get`:

```sh
$ go get github.com/SethMMorton/secret-santa
```

You will find `secret-santa` installed in `$GOPATH/bin`.

# Example

Suppose that you want to create a Secret Santa for a group of friends.
You would create a YAML where each participant is an element in the
file. The person's name appears under the `name` attribute. For example:

```yaml
- name: Steve
- name: Jane
- name: Karthik
- name: Molly
- name: Patrick
- name: Mel
- name: Peggy
- name: Jim
- name: Mildred
- name: Carl
- name: Dr. Bob
```

If you execute `secret-santa` on a file with the above contents, it will
assign recipients for each participant, ensuring there are no self-recipients
and no circular pairs (A gives to B, B gives to A).

However, in practice, it is often desirable that people that are very close
(e.g. spouses, significant others, etc.) not get each other in the Secret
Santa because they were already buying each other gifts. To solve this problem
you can assign a list of "invalid" recipients for each partipant. The
`secret-santa` program will not allow these people to be recipients for that
participant. Here is how that might look:

```yaml
- name: Steve
  invalid:
    - Jane
    - Karthik
- name: Jane
  invalid:
    - Steve
    - Karthik
- name: Karthik
  invalid:
    - Steve
    - Jane
- name: Molly
  invalid:
    - Patrick
- name: Patrick
  invalid:
    - Molly
- name: Mel
  invalid:
    - Peggy
- name: Peggy
  invalid:
    - Mel
- name: Jim
  invalid:
    - Mildred
- name: Mildred
  invalid:
    - Jim
- name: Carl
- name: Dr. Bob
```

To determine who is buying for whom, simply execute the `secret-santa` program.
Here is what you might get:

```sh
$ secret-santa sample.yaml
Peggy   ==> Carl
Mildred ==> Patrick
Carl    ==> Mel
Steve   ==> Peggy
Jane    ==> Mildred
Karthik ==> Jim
Jim     ==> Steve
Dr. Bob ==> Jane
Molly   ==> Dr. Bob
Patrick ==> Karthik
Mel     ==> Molly
```

Each time you run the program the results will be different, so don't expect to
get the above results should you use this exact example.
