Ext.define('ToDo.view.toDoList.ToDoList', {
    extend: 'Ext.panel.Panel',

    /* Marks these are required classes to be to loaded before loading this view */
    requires: [
        'ToDo.view.toDoList.ToDoListController',
        'ToDo.view.toDoList.ToDoListModel'
    ],
       
    xtype: 'app-todoList',
    controller: 'todoList',

    /* View model of the view */

    viewModel: {
        type: 'todoList'
    },

    items: [{
        xype: 'container',
        items: [ 
        {
            xtype: 'container',
            layout: 'hbox',
            cls: 'task-entry-panel',
            defaults: {
                flex: 1
            },
            items: [
                {
                    reference: 'newToDo',
                    xtype: 'textfield',
                    emptyText: 'Enter a new todo here'
                },
                {
                    xtype: 'button',
                    name: 'addNewToDo',
                    cls: 'btn-orange',
                    text: 'Add',
                    maxWidth: 50,
                    handler: 'onAddToDo'
                }]
            }
        ]
    }]
});