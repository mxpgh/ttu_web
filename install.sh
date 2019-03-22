#!/bin/sh

basepath=$(cd `dirname $0`; pwd)
echo "current path: ${basepath}"

install_ttu_web() {
	cp $basepath/libsysconfig.so $basepath/libyaml-cpp.so $basepath/libyaml-cpp.so.0.6 /usr/lib
	cp $basepath/sysConf.yaml /etc	

	RCLOCAL="/etc/rc.local"
	if test -f ${RCLOCAL}
	then
		cp ${RCLOCAL} "${RCLOCAL}.bak"
		if grep -q $basepath/ttu_web "${RCLOCAL}.bak";then
			echo ""
		else
			sed -i "/^[^#]*exit/ i nohup $basepath/ttu_web > $basepath/ttu_web.log 2>&1 &" ${RCLOCAL}
		fi
		rm -rf "${RCLOCAL}.bak"
	fi

	chmod +x ${RCLOCAL}
	chmod a+x $basepath/start.sh
	/bin/sh -c $basepath/start.sh
}

install_ttu_web

