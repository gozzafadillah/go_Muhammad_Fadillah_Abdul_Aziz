#!/bin/sh
days=$(date)
echo "Nama : "
	read nama
echo "FB : "
	read fb
echo "Linkedin : "
	read linkedin
nameDir="${nama} ${days}"
pingGoogle="$(ping -c 3 google.com)"
mkdir "${nameDir}"

host=$(uname -a)
username="$USER"

cd "${nameDir}"
mkdir about_me 
mkdir my_friend 
mkdir my_system_info

cd about_me
mkdir personal profesional
cd personal
cat << EOF > facebook.txt
https://www.facebook.com/${fb}
EOF
cd ..
cd profesional
cat << EOF > linkedin.txt
"https://www.linkedin.com/in/${linkedin}"
EOF

cd ../../my_friend

cat <<EOF > list_of_my_friend.text
1) Achmad Miftahul - amulum
2) Achmad Rafiq - achmadrafiq97
3) Adiststi - adiststi
4) Agung - agungajin19
5) Azzahra - al7262
6) Charisma - fadzrichar
7) Farida - ulfarida
8 ) Garry - garryarielcussoy
9) Hamdi - hamdiranu
10) Hedy Gading - Gading09
11) Ilham - aamfatur
12) Lelianto - Lelianto
13) Mohammad - daffa99
14) Muhammad Fadhil - fabdulkarim
15) Ofbimon - bimon-alta
16) Purnama Syafitri - pipitmnr
17) Setyo - setyoyo07
18) Wildan - wiflash
19) Willy - sumarnowilly94
20) Woka - woka20
EOF

cd ../my_system_info

cat <<EOF > about_this_laptop.txt
My username: ${username}
With host: ${host}
EOF
cat <<EOF > internet_connection.txt
${pingGoogle}
EOF