# demos
[Tools][6] for topological line simplification in the context of arbitrary
planar objects (point, lines, and polygons).  

## constrained RDP algorithm
Contextual simplification of polylines using the 
[RDP Algorithm][0]. The [algorithm][1] is a recursive decomposition of 
a polyline given the maximum offset distance threshold (![epsilon][2]) 
of the simplified(![lprime][4]) from the original (![lprime][3]).

[Read more ...][5]

[Download application ...][6]

[0]: <https://utpjournals.press/doi/10.3138/FM57-6770-U75U-7727> "RDP"
[1]: <https://en.wikipedia.org/wiki/Ramer%E2%80%93Douglas%E2%80%93Peucker_algorithm> "RDP Wiki"
[2]: <https://latex.codecogs.com/svg.latex?%5Cinline%20%5Clarge%20%5Cvarepsilon> "\varepsilon"
[3]: <https://latex.codecogs.com/svg.latex?%5Cinline%20%5Clarge%20L> "L"
[4]: <https://latex.codecogs.com/svg.latex?%5Cinline%20%5Clarge%20L%5E%5Cprime> "L^\prime"
[5]: <https://github.com/TopoSimplify/demos/blob/master/constrainedRDP/README.md> "demos RDP"
[6]: <https://github.com/TopoSimplify/demos/tree/master/dist/constdp> "download App"
