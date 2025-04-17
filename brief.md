# Submission Hypermedia demo

The aim of this project is to demonstrate how using hypermedia with HTMX can make re-usable and cohesive components that are simpler than the ad-hoc approaches we're taking currently on a project.

## Some domain info

### Submission

A submission refers to a submission to a peer review system. In real life, there are a large number of fields, but for the purpose of this demo they have

- Title
- Abstract
- Keywords
- Author

### Submission controls

An editor, at any point, should be able to do the following to a submission:
- Approve
- Reject
- Add to watchlist

I would expect our submission data type to hold all of this information, so that it can be rendered.

## Requirements

This project should have a number of dummy pages that can be navigated to using a menu.

Inside them, there can be a bunch of dummy forms that an editor can be filling out to do work. What they do is not important, but it would be good if it seemed semi-plausible. 

In each page, it should have a visual representation of the submission the editor is working on. This should be consistent on all the pages. This is where hypermedia and HTMX comes into play. I expect each page to use HTMX to pull in the hypermedia representation, along with the controls to do things like approve, reject, and add to watchlist. It should also mean it looks the same on all pages.

## Intent

This should show the power of having a single hypermedia representation that can be used to do all the things to a submission. By using HTMX we also don't have to worry about redirecting to the specific page the user was on, because the control is the only thing that would reload. We can also demonstrate that we can evolve the submission hypermedia with new controls and then they should appear on all pages for free. 

## Shortcuts for the demo

- Just have a canned submission that the user can perform the actions on. Don't worry about persistence
- I'm also not worried about tests in this demo. 
- With respect to the dummy pages, 3 is enough

## Tech choices

- Go
- HTMX

Do not bring in anything else. Vanilla JS and CSS is fine. It doesn't need to look amazing. 