## Constrained RDP Algorithm
Constrained simplification of arbitrary polylines in the context 
of arbitrary planar geometries.

### How to use 
Open a terminal (command line) from the directory containing constdp[.exe]. Simplification options are made available through  the use of [TOML][0] file (config.toml). Execute constdp with the following command :
```bash
./constdp -c ./config.toml 
```
If a `-c` option is not provided at the terminal e.g. `./constdp`, it assumes `-c ./config.toml` as default.

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
#### Data

#### Constraints


[0]: <https://github.com/toml-lang/toml> "TOML"
