---
layout: default
title: Go Cookbook
---

{% for chapter in site.data.chapters %}
  <h3> {{ chapter.title }}</h3>

  <ul>
  {% for recipe in chapter.recipes %}
    <li>
    {% if recipe.wip %}
      <strong>[work in progress]</strong>
    {% endif %}
    {% if recipe.path %}
      <a href="{{recipe.path}}">{{ recipe.title }}</a>
    {% else %}
      {{ recipe.title }}
    {% endif %}
    </li>
  {% endfor %}
  </ul>
{% endfor %}

