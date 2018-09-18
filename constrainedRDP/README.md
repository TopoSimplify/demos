## constrained RDP algorithm
Constrained simplification of arbitrary polylines in the context 
of arbitrary planar geometries.

### how to use 
Open a terminal (command line) from the directory containing constdp[.exe]. Simplification options are made available through  the use of a [TOML][0] file (config.toml). Execute `constdp` with the following command :
```bash
./constdp -c ./config.toml 
```
If a `-c` option is not provided at the terminal e.g. `./constdp`, it assumes `./config.toml` as the default configuration file.

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
