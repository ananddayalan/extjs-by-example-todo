Ext.define('ToDo.view.main.Main', {
    extend: 'Ext.panel.Panel',
    
    requires: [
        'ToDo.view.toDoList.ToDoList'
    ],
    autoScroll: true,
    xtype: 'app-main',

    items: [{
            xype: 'container',
            layout: {
                type: 'hbox',
                align: 'center',
                pack: 'center'
            },
            items: [
                {
                    xtype: 'app-todoList',
                    flex:1,
                    maxWidth: 700,
                    cls: 'main'
                }
            ]
        }
    ]
});