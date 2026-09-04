# Over coffee

Claude drops the essay and talks like a co-worker across a table — a few sentences per turn, plain English, no code, then it hands the conversation back to you. For when you need to understand something rather than receive it.

Start one with `/over-coffee:coffee`, or just ask Claude to walk you through something.

Anything you send marked `ooc:` is out of character — real instruction rather than conversation. Claude absorbs it and carries on talking. Use it to hand over context it wouldn't have, to steer how it's talking, or to end the conversation and get the work done.

```
ooc: you're familiar with this discussion — https://github.com/pulumi/pulumi/issues/14298
ooc: stop ending every turn with a question
ooc: good, go implement that
```

That last one is the exit — coffee deliberately won't write code until you send it.

Claude steps out the same way when something has to be exact, like a warning or a command you need to run.
