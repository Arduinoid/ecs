# Entity Component System

This is a experiment in order to better understand how an ECS based system
works. I hope to use this as an inspiration for new and interesting ways to
organize code. The interface inference in Go makes this a fun little
implementation as well. At some point I may try to use this along side a
graphics lib such as [SDL](https://github.com/veandco/go-sdl2) or
[Raylib](https://github.com/gen2brain/raylib-go)

This ECS is heavily inspired by [Engo's ECS](https://github.com/EngoEngine/ecs)
which is used in it's own game library
[Engo](https://github.com/EngoEngine/engo)

## Design

My design goal for this is to have entities be automatically registered to any
system that they qualify for based on their component make up. This is achieved
using interfaces to help identify the components in an entity upon adding them
to the world. The system does not yet support dynamic component adding or
removing on entities but this is in the works.

## Why ECS?

**What advantages do you get when using the ECS pattern and why?**

Here are just a few top picks that come up often to answer the question:

- Composability of entities (otherwise objects)
- Better code reuse
- Entities can quickly adopt new behavior with little change
- In some cases better performance

## Composability

This is done by taking our traditional objects that would contain both _data_
and _behavior_ and breaking them up into _components_ (data) and _systems_
(behavior) then reassembling or relating them back together through an entity.
The entity can either be a wrapper or simply an ID. Entities only hold
components and do not act on them. It is the job of the system to act on the
components for each of the entities.

Components themselves are just data containers for example a position component

```go
type Position struct {
    X, Y int
}
```

Or something a bit more complex like a image

```go
type Image struct {
    Width, Height int
    Path string
    Scale float32
}
```

Providing simple data components allows you to build up entities using existing
work which leads to the next point

## Reusability

As you build new components, because of there single and usually small purpose,
they can be used over and again. This makes composing new entity types much
faster and quite a bit easier too. Since the behavior is no longer part of the
entity or the components, it makes it easier to build a new component without
out a lot of cascading changes.

## Easily Extend Behavior

Since our behavior is centralized in various systems we only need to add the
components that a system acts on in order to extend entity behavior. Say we want
our entity to have movement where it can be positioned in space but have
velocity so it can change positions over time. If there was a system for
movement and it only requires that an entity have a _Postion_ component and a
_Velocity_ component then it would be able to act on that entities data. If the
entity contained other components as well, other related systems would be able
to provide other behaviors like gravity, player input, rendering, and so on.

## Performance

The performance gains comes for the idea of "data locality" and does come down
to the implementation but the idea is when you have all of your data together,
say a list of components, and you are looping over and changing that data, it is
a lot faster for the CPU if all of it is sitting together in contiguous memory.
More on this later, to be continued...
