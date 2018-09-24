## constrained RDP algorithm
Constrained simplification of arbitrary polylines in the context 
of arbitrary planar geometries.

### how to use 
Open a terminal (command line) from the directory containing an executable (constdp[.exe] for 64bit, constdp_32bit[.exe] for 32bit systems). Simplification options are made available through  the use of a [TOML][0] file (config.toml). Execute `constdp` with the following command :

```bash
./constdp -c ./config.toml 
```

If a `-c` option is not provided at the terminal e.g. `./constdp`, it assumes `./config.toml` as the default configuration file.
Change `config.toml` to configure your simplification. 

#### config file 

```toml
#input file is required
Input                  = "/path/to/input.[wkt]" 
#output is optional, defaults to ./out.txt
Output                 = "" 
#this is optional
Constraints            = "/path/to/file.[wkt]" 
#options : DP, SED
SimplificationType     = "DP" 
Threshold              = 0.0
MinDist                = 0.0
RelaxDist              = 0.0
#are polylines independent or a feature class ?
#if false planar and non-planar intersections
#between polylines are not observed
IsFeatureClass         = false
#observe planar self-intersection
PlanarSelf             = false
#observe non-planar self-intersection
NonPlanarSelf          = false
#avoid introducing new self-intersections as a
#result of simplification
AvoidNewSelfIntersects = false
GeomRelation           = false
DistRelation           = false
SideRelation           = false
```

### data 
Input in `config.toml` should point to a text file containing [WKT][4]  strings or `toml` arrays. 

#### wkt input
```bash
LINESTRING (30 10, 10 30, 40 40)
# linestring with 3d coordinates
LINESTRING (30 10 1, 10 30 2, 40 40 3)
```
See sample input and constraints WKT text files : [Input][7], [Constraints][8].

#### toml input
```toml
1=[[30, 10], [10, 30], [40, 40]]
2=[[30,  8], [10, 15], [40, 25]]
#lines with 3d
3=[[30.1, 8.2, 2.4], [10.4, 15.9, 5.6], [40.8, 25.0, 9.8]]
```


Note that the `toml` input uses an `id=array`, contents of the array must be of the same type ( all coordinates as 
integers or floats). A point is `[x , y]` or `[x, y, z]`. A polyline is a string of points `[[x,y],[x,y],...]`. A polygon is a string of of polylines `[[[x,y],[x,y],...]]`; the fist is a shell and subsequent strings are holes e.g.

```toml
1=[[[35, 10],[45, 45],[15, 40],[10, 20],[35, 10]], [[20, 30],[35, 35],[30, 20],[20, 30]]]
```

See sample input and constraints `toml` text files : [Input][5], [Constraints][6].
Since constraints can be of the form `point, polylines, or polygon` its `toml` is of the format:

```bash
[points]
id=array
id=array

[polylines]
id=array 
id=array 

[polygons]
id=array 
id=array 
```

#### example
Given a polyline in `resource/input.wkt`

```toml
Input                  = "resource/input.wkt"
Output                 = ""
Constraints            = "resource/constraints.wkt"
SimplificationType     = "DP"
Threshold              = 50.0
MinDist                = 20.0
RelaxDist              = 30.0
IsFeatureClass         = false
PlanarSelf             = true
NonPlanarSelf          = true
AvoidNewSelfIntersects = true
GeomRelation           = true
DistRelation           = true
SideRelation           = true
```
Original polyline in the context of planar objects: 
![polyline][1]

Constrained simplification with respect to config options(above): 
![polyline][2]

Unconstrained simplification with these options turned `false`:

```toml
IsFeatureClass         = false
PlanarSelf             = false
NonPlanarSelf          = false
AvoidNewSelfIntersects = false
GeomRelation           = false
DistRelation           = false
SideRelation           = false
```
![polyline][3]


[0]: <https://github.com/toml-lang/toml> "TOML"
[1]: <./resource/original.png> "original Polyline"
[2]: <./resource/simple1.png> "simple 1"
[3]: <./resource/simple2.png> "simple 2"
[4]: <https://en.wikipedia.org/wiki/Well-known_text> "wkt wiki"
[5]: <./resource/input.toml> "input toml"
[6]: <./resource/constraints.toml> "constraints toml"
[7]: <./resource/input.wkt> "input wkt"
[8]: <./resource/constraints.wkt> "constraints wkt"
