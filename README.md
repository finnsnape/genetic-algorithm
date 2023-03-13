# Genetic Algorithm
## Potential libraries
- [goga](https://github.com/tomcraven/goga) -- GA (golang)
- [box2d](https://github.com/ByteArena/box2d) -- physics (golang)
- [oak](https://github.com/oakmound/oak) -- physics (golang)
- [pixel](https://github.com/faiface/pixel) -- physics (golang)
- [engo](https://github.com/EngoEngine/engo) -- physics (golang)
- [wails](https://github.com/wailsapp/wails) -- GUI (golang)
- [openga](https://github.com/Arash-codedev/openGA) -- GA (c++)
- [box2d](https://github.com/erincatto/box2d) -- physics (c++)

## TODO
- Add comments
- UI (take advantage of svelte)
- link UI to GA with wails
- Goroutines
- Simplify CSS (won't need to be responsive)
- elitism
- Reduce code duplication

## Notes
- Investigate using different goals (e.g., highest velocity, furthest distance)
- Maybe use a central leaderboard for storing which student got the best result
- Allow people to export configurations and/or generations
- Maybe don't show every generation
- Important to be able to tell what is and isn't a good algorithm. Perhaps have a "bad" preset and an "okay" preset
- Allow them to write some of the functions themselves
- Perhaps a less advanced mode for GCSE/A-Level students. Bundled as .exe?
- Abstract away all the graphical stuff so the students can investigate just the GA code
- Ensure basic preset doesn't take too long to run
- Electron-esque window for config, SDL window for visuals?
- Graphs
- Jupyter style
- Abstract away stuff that they shouldn't be able to edit
- User stories from perspective of professor, not students. Performance is relevant to have it work within the time frame, and features are useful for learning outcomes
- Note down choices about libraries
- Create worksheets
- Competition?
- code security