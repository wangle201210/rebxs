# rebxs
忍者必须死-疯人院-积分-悬赏

### 表格分析
#### 1.人员表 （user）
	name: 人员名字

#### 2.项目表（project）
	name: 项目名字

#### 3.项目对应积分算法表（score）
	project_id: 对应项目
	min: 最小值
	max: 最大值
	score: 积多少分

	eg1: 项目表里添加一个项目名为河豚（id=1）,score表里面添加一条记录project_id=1 min=109 max=999 score = 4 
		意思即为河豚副本大于等于109小于等于999则积4分

	eg2: 项目表里添加一个项目名为金币（id=2）,score表里面添加一条记录project_id=2 min=1600 max=1639 score = 2 
	
	eg3: 项目表里添加一个项目名为主将（id=3）,score表里面添加一条记录project_id=3 min=1 max=1 score = 3

	eg4: 项目表里添加一个项目名为团战一（id=4）,score表里面添加一条记录project_id=4 min=2 max=5 score = 2

	eg5: 项目表里添加一个项目名为家族贡献（id=5）,score表里面添加一条记录project_id=5 min=1 max=1 score = 3
		意思即为家族贡献为1（此贡献非直接计算家族贡献积分的贡献值，此处想表达的是比如1就是管理员，2就是族长这类）此处为1则理解为管理员，则积3分
	（以上是我能想到的所有积分出现的情况，若还有未考虑到的，请批评指正）

#### 4.悬赏表（reward）
	目前就直接按照固定算法，即s记1分，s+记2分，ss和ss+记4分
	user_id:提供者
	type:类型（s|s+|ss）
	participant:参与人员
	time:时间

#### 5.记录表（record）
	user_id: 对应人员数据
	project_id: 对应项目
	result: 得分
	time: 时间（时间考虑是自动根据添加时计算，还是手动录入，若自动计算，则管理员要实时录入，比如周天的家族战数据，若下周一才录入将会计入下一周的积分，这样或许不妥，如果手动录入会加大录入人员工作量）

#### 6.汇总表，此表仅展示不储存（summary）
	根据现在统计方式计算汇总表需要显示内容为
	a.时间
	b.人员信息
	c.积分信息 积分信息会分为多列显示，会获取当前时间所有人的所有积分表，然后按照项目分组后得出表头
	d.排名

	操作方式为：先选择一个时间（段），然后选择查看方式（分积分和成绩两种），最后点击查看，将会显示在页面，如果需要可以导出成excel
	最终生成表样式为
	| 成员 | 项目 ** | 项目 ** | ... | 悬赏 | 总积分 | 排名 |
	| 小张 |   2    |   3     | ... | 5   |  18    |  3  |


### 可行性（相较于直接用excel）
    当把基础设置项录入（或者我直接设置为固定数据）后，管理员只需要把每个人的成绩录入即可，不需要再去对照积分表算出没人每项积分，然后再去汇总（当然excel有自动汇总的功能）
	如果用excel每次要给工作表设置公式，即使只是简单的求和和排序也需要时间
	方便展示，不用每次都上传excel到群，管理员直接也不用相互等数据（比如上周出现的问题：柯基把数据录好了，结果忘记发给其他管理员，其他人就只能等着柯基睡醒了才能安排悬赏）
	方便扩展，比如后续如果做半年汇总，我可以提供原始数据、再如若呗安排进悬赏后相应要扣除一部分分值，比如被安排s后，对应总积分减一（方便后续的人被安排，不然每周基本都是那几位才能打到悬赏）
	last but not least 高端大气上档次


### 缺点 
    程序的初始版本 2、3表将会定下来，不给修改功能，若需修改需要联系我（之后可能会添加修改积分规则方法）
	实质上还是绕不过录入人员数据的阶段，并且录入数据还没有excel那么方便


### 其他
    网站服务器和域名可以用我的，我的服务器2021年才到期
	开发上，如果感觉可行，我就当练习下我之前学习的新语言（golang）的写法，顺便为家族做点贡献
	担心的就是，我忙了几天，结果又不用～～这样很伤人的。


