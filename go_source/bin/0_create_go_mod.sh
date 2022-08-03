
bin_dir=`pwd`
cd $bin_dir
cd ../
project_path=`pwd`

echo "========================================================"
echo "Project Path >> "
echo ": "$project_path
echo "========================================================"
echo "Clean module >>"
echo "file : " $project_path/src/go.mod
rm -rf  $project_path/src/go.mod
echo "========================================================"
echo "Create module init(kbell) >>"
cd $project_path/src
go mod init kbell
echo "========================================================"
echo "module show >>"
echo "file : " $project_path/src/go.mod
cat $project_path/src/go.mod
echo "========================================================"

cd $bin_dir
./0_get_gopkgs.sh
