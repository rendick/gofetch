#!/bin/bash
copy () {
	chmod +x gofetch
	if [[ $1 == "yes" ]] ; then
		sudo chmod +x gofetch
		sudo mv gofetch /bin/
	elif [[ $1 == "no" ]] ; then
		exit
		else
		echo 'invalid installation type'
	fi
}

PS3='Installation type: '
types=("yes" "no")

echo "Note, that you need to install xrands and lsb-release!"

select type in "${types[@]}" ; do
	if [[ -n $types ]] ; then
		echo 'You selected: $1'
		break
	else
		echo 'Please select between user and system installation'
	fi
done

sleep 2

copy $type
