# TODO

- Finish pause menu
- Break up the cell behavior into components:
    - movable
    - blocking 
    - stackable?
    - this would require pulling out the sticky translation into its own system that handles the collision and rotation
    - this system would need to get all cells that are movable and all entries that are not movable and blocking. these are the ones to check collision for.
- 