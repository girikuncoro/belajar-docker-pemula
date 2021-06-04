<template>
  <div class="todo pa-6">
    <v-text-field
      v-model="newTaskTitle"
      @click:append="addTask"
      @keyup.enter="addTask"
      class="pa-3"
      outlined
      label="Add Task"
      append-icon="mdi-plus"
      hide-details
      clearable
    ></v-text-field>
    <v-list class="pt-0" flat>
      <div v-for="task in tasks" :key="task.id">
        <v-list-item
          @click="doneTask(task.id)"
          :class="{ 'blue lighten-5': task.done }"
        >
          <template v-slot:default>
            <v-list-item-action>
              <v-checkbox :input-value="task.done" color="primary"></v-checkbox>
            </v-list-item-action>

            <v-list-item-content>
              <v-list-item-title
                :class="{ 'text-decoration-line-through': task.done }"
              >
                {{ task.note }}
              </v-list-item-title>
            </v-list-item-content>
            <v-list-item-action>
              <v-btn @click.stop="deleteTask(task.id)" icon>
                <v-icon color="grey lighten-1">mdi-trash-can</v-icon>
              </v-btn>
            </v-list-item-action>
          </template>
        </v-list-item>
        <v-divider></v-divider>
      </div>
    </v-list>
  </div>
</template>

<script>
import axios from "../libs/axios";

export default {
  name: "Todo",
  data() {
    return {
      newTaskTitle: "",
      tasks: [],
    };
  },
  async mounted() {
    await this.fetchTasks();
  },
  methods: {
    async fetchTasks() {
      try {
        const response = await axios.get("/todo");
        this.tasks = response.data;
      } catch (error) {
        console.dir(error);
      }
    },
    async addTask() {
      let newTask = {
        note: this.newTaskTitle,
        done: false,
      };
      this.tasks.push(newTask);
      this.newTaskTitle = "";

      try {
        await axios.post("/todo", newTask);
      } catch (error) {
        console.log(error);
      }
    },
    async doneTask(id) {
      let task = this.tasks.filter((task) => task.id == id)[0];
      task.done = !task.done;

      try {
        await axios.put("/todo", task);
      } catch (error) {
        console.log(error);
      }
    },
    async deleteTask(id) {
      this.tasks = this.tasks.filter((task) => task.id !== id);

      try {
        await axios.delete("/todo", { data: { id: id } });
      } catch (error) {
        console.log(error);
      }
    },
  },
};
</script>
