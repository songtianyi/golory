#!/bin/bash

#!/bin/bash
myheader="// Copyright 2018 golory Authors @1pb.club. All Rights Reserved.\n\
//\n\
// Licensed under the Apache License, Version 2.0 (the \"License\");\n\
// you may not use this file except in compliance with the License.\n\
// You may obtain a copy of the License at\n\
//\n\
//      http://www.apache.org/licenses/LICENSE-2.0\n\
//\n\
// Unless required by applicable law or agreed to in writing, software\n\
// distributed under the License is distributed on an \"AS IS\" BASIS,\n\
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n\
// See the License for the specific language governing permissions and\n\
// limitations under the License.\n"

set -o errexit -o nounset
# check format
# check license header
find . -name \*.go -exec sh -c "if ! grep -q 'LICENSE-2.0' '{}';then mv '{}' tmp && echo '$myheader' > '{}' && cat tmp >> '{}' && rm tmp;fi" \;