#ifndef __TTU_SYS_CONFIG__
#define __TTU_SYS_CONFIG__

#ifdef __cplusplus
extern "C" {
#endif 

//返回cpu占用率，百分比.此函数运行时间1秒钟
extern int getCpuOccupy();

//返回内存大小,单位KB
extern int getRamSize();

//获取磁盘大小，单位GB
extern float getDiskSize();

//返回内存占用率，百分比
extern int getRamOccupy();

//调用命令dh获取磁盘信息,返回磁盘占用百分比
extern int getDiskOccupy();

//获取设备运行时间,返回秒
extern unsigned long getRunTime();

//获取启动时间，格式 2019-01-02 10:43:57.失败返回0.成功返回字符串长度
extern int getUPTime(char * buffer,int buffersize);

//获取设备当前时间，格式 2019-01-02 10:43:57.失败返回0.成功返回字符串长度
extern int getTime(char * buffer,int buffersize);

//设置设备当前时间，格式 2019-01-02 10:43:57
extern int setTime(char * time);

//获取，设备类型.失败返回0.成功返回字符串长度
extern int getDevType(char * buffer,int buffersize);

//获取，设备名称.失败返回0.成功返回字符串长度
extern int getDevName(char * buffer,int buffersize);

//获取，设备状态.失败返回0.成功返回字符串长度
extern int getDevStatus(char * buffer,int buffersize);

//获取，设备厂商信息.失败返回0.成功返回字符串长度
extern int getDevVendor(char * buffer,int buffersize);

//获取，硬件版本.失败返回0.成功返回字符串长度
extern int getHardwareVer(char * buffer,int buffersize);

//获取，平台软件版本.失败返回0.成功返回字符串长度
extern int getSoftwareVer(char * buffer,int buffersize);

//获取RTC状态 错误返回0,正常返回1
extern int getRtcStatus();

//获取设备温度。成功返回1，传感器故障返回0
extern int getTemperature(int * t);


//设置温度报警阈值,成功返回1，失败返回0.
extern int setTempThreshold(int upper,int lower);
//获取温度报警阈值,成功返回1，失败返回0
extern int getTempThreshold(int * upper,int * lower);
//获取温度报警监控判定窗口时间
extern int getTempInterval(int * value);
//设置温度报警监控判定窗口时间
extern int setTempInterval(int value);


//获取系统监控窗口时间。监控运行间隔时间，单位分钟。成功返回1，失败返回0
extern int getAlarmInterval(int * interval);
//设置系统监控窗口时间。监控运行间隔时间，单位分钟。成功返回1，失败返回0
extern int setAlarmInterval(int     interval);
//设置cpu报警阈值(百分比),成功返回1，失败返回0
extern int setCpuThreshold(int upper);
//获取cpu报警阈值(百分比),成功返回1，失败返回0
extern int getCpuThreshold(int* upper);
//设置ram报警阈值(百分比),成功返回1，失败返回0
extern int setRamThreshold(int upper);
//获取ram报警阈值(百分比),成功返回1，失败返回0
extern int getRamThreshold(int* upper);
//设置disk报警阈值(百分比),成功返回1，失败返回0
extern int setDiskThreshold(int upper);
//获取ram报警阈值(百分比),成功返回1，失败返回0
extern int getDiskThreshold(int* upper);

//设置主站ip port,成功返回1，失败返回0
//main = 0	主站
//main = 1	备用主站
//ver = 4	ipv4
//ver = 6	ipv6 
extern int setServer(int main,int ver,char * ip,int port);

//获取主站ip port,成功返回1，失败返回0
//main = 0	主站
//main = 1	备用主站
//ver = 4	ipv4
//ver = 6	ipv6 
extern int getServer(int main,int ver,char * ip,int * port);


typedef enum _PARAMETER{
	cpu,		//cpu使用率上限
	ram,		//ram使用率上限
	interval	//监控判定窗口时间
}PARAMETER;
typedef enum _CATEGORY{
	docker,
	app
}CATEGORY;
//获取docker，app监控参数。成功返回1，失败返回0
extern int getMonParameter(CATEGORY category,PARAMETER parameter,int * value);
//设置docker，app监控参数。成功返回1，失败返回0
extern int setMonParameter(CATEGORY category,PARAMETER parameter,int value);




//系统重启
extern int reboot();


#ifdef __cplusplus
}
#endif 

#endif

