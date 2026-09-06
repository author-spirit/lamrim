const workflowsEl = document.getElementById("workflows");
const outputEl = document.getElementById("output");

async function loadWorkflows() {
  const response = await fetch("/api/workflows");
  if (!response.ok) {
    throw new Error(await response.text());
  }
  return response.json();
}

function renderWorkflows(workflows) {
  workflowsEl.innerHTML = "";

  if (workflows.length === 0) {
    workflowsEl.textContent = "No workflows found.";
    return;
  }

  for (const workflow of workflows) {
    const card = document.createElement("article");
    card.className = "card";

    const title = document.createElement("h2");
    title.textContent = workflow.name;

    const description = document.createElement("p");
    description.textContent = workflow.description || "No description.";

    const button = document.createElement("button");
    button.textContent = "Run";
    button.addEventListener("click", () => runWorkflow(workflow.name));

    card.append(title, description, button);
    workflowsEl.append(card);
  }
}

async function runWorkflow(name) {
  outputEl.textContent = `Running ${name}...`;

  const response = await fetch(`/api/workflows/${encodeURIComponent(name)}/run`, {
    method: "POST",
  });

  const body = await response.json();
  outputEl.textContent = JSON.stringify(body, null, 2);

  if (!response.ok) {
    throw new Error(body.error || "run failed");
  }
}

loadWorkflows()
  .then(renderWorkflows)
  .catch((error) => {
    outputEl.textContent = error.message;
  });
