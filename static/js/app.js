async function fetchTasks(){
const res = await fetch('/api/tasks');
const data = await res.json();
return data;
}


function renderTasks(tasks){
const ul = document.getElementById('tasks');
ul.innerHTML = '';
if(tasks.length === 0){ ul.innerHTML = '<li>No tasks yet.</li>'; return }
for(const t of tasks){
const li = document.createElement('li');
li.dataset.id = t.id;
li.innerHTML = `<label><input type="checkbox" class="toggle" ${t.done? 'checked' : ''} /> <span class="title">${t.title}</span> <small class="meta">${new Date(t.created_at).toLocaleString()}</small></label>`;
ul.appendChild(li);
}
}


async function load(){
const tasks = await fetchTasks();
renderTasks(tasks);
}


async function addTask(title){
await fetch('/api/tasks', { method:'POST', headers:{'Content-Type':'application/json'}, body: JSON.stringify({ title }) });
await load();
}


async function toggleTask(id){
await fetch(`/api/tasks/${id}/toggle`, { method:'POST' });
await load();
}


document.addEventListener('submit', function(e){
if(e.target && e.target.id === 'addTaskForm'){
e.preventDefault();
const input = document.getElementById('title');
const v = input.value.trim();
if(!v) return;
addTask(v);
input.value = '';
}
});


document.addEventListener('click', function(e){
if(e.target && e.target.classList.contains('toggle')){
const li = e.target.closest('li');
const id = li && li.dataset && li.dataset.id;
if(id) toggleTask(id);
}
});


window.onload = load;